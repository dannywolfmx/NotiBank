const electron = require('electron');
const path = require('path');
const request = require('tinyreq');
const notifier = require('node-notifier');

const app = electron.app;
const Tray = electron.Tray;

//Definicion de los iconos usados
const iconPath = path.join(__dirname, 'app-icon.png');
const arrowDown = path.join(__dirname, 'down-green.png');
const arrowUp = path.join(__dirname, 'up-red.png');
const equalIcon = path.join(__dirname, 'equal.png');

// Module to control application life.
const SONIDO_WINDOWS = 'Notification.Reminder';
const tiempo = 1800000; // 30 minutos
const urlBanco =
  'https://www.banorte.com/wps/portal/ixe-xima/Home/indicadores/tipo-de-cambio/!ut/p/a1/04_Sj9CPykssy0xPLMnMz0vMAfGjzOLdjQwtPIydDbz9vRxdDRwd_TwtHAP8DFyNTYEKIvEoCDAhTr-zu6OHibmPgYG_ibuRgaOFn5dJsKmlkYGnGXH6DXAARwNC-sP1o_AqAfkArACfE8EK8LihIDc0NMIg0xMAQCluTg!!/dl5/d5/L2dBISEvZ0FBIS9nQSEh/';

let tray = null;

//La aplicacion por el momento solo guarda un antecedente de tipo de cambio, mientras esta en ejecucion
//Una vez cerrado el programa este dato se cierra
var tipoDeCambioAnterior = -1;

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow;

function createWindow() {
  //Boton que aparece en la barra de herramientas
  tray = new Tray(iconPath);

  tray.setToolTip('Notificar tipo de cambio');

  tray.on('click', () => {
    notificacionTipoCambio();
  });

  tray.setToolTip('Actualizar estado');

  notificacionTipoCambio();
  setInterval(notificacionTipoCambio, tiempo);
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow);

// Quit when all windows are closed.
app.on('window-all-closed', function() {
  // On OS X it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('activate', function() {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (mainWindow === null) {
    createWindow();
  }
});

function bufferToString(x) {
  return x.toString('utf8');
}

function notificacion(titulo, texto, icono, sonido) {
  const mensaje = {
    title: 'Tipo de cambio actual',
    message: texto,
    icon: icono, // Absolute path (doesn't work on balloons)
    sound: false, //sonido, // Only Notification Center or Windows Toasters
    wait: true, // Wait with callback, until user action is taken against notification,
    appId: 'Hola',
  };

  notifier.notify(mensaje, function(err, response) {
    // Response is response from notification
  });
}

function dameTipoDeCambio(texto) {
  //Palabra clave para encontrar el tipo de cambio
  let cadenaReferencia = '"nombreDolar":"VENTANILLA"';
  let posicionInicial = 53; // Ejem: [{"nombreDolar":"VENTANILLA","compra":"18.05","venta":"19.45"}]}' // son 53 saltos para llegar al cambio de venta
  let posicionFinal = 5; // Ejem: "19.04" // 5 caracteres

  let posicion = texto.search(cadenaReferencia);
  return texto.slice(
    posicion + posicionInicial,
    posicion + posicionInicial + posicionFinal,
  );
}

function notificacionTipoCambio() {
  //Tipo de cambio banorte
  request(urlBanco, function(err, body) {
    if (err) {
      console.log(err);
    } else {
      let tipoCambio = dameTipoDeCambio(body);
      let icono = null;
      let sonido = false;

      if (tipoDeCambioAnterior < tipoCambio) {
        icono = arrowUp;
        sonido = SONIDO_WINDOWS; //redefinir variable de boolean a string
      } else if (tipoDeCambioAnterior == tipoCambio) {
        icono = equalIcon;
      } else if (tipoDeCambioAnterior > tipoCambio) {
        //Verificar que si sea numerico
        icono = arrowDown;
        sonido = SONIDO_WINDOWS; //redefinir variable de boolean a string
      }

      if (icono) {
        notificacion('Tipo de cambio', tipoCambio, icono, sonido);
        tipoDeCambioAnterior = tipoCambio;
      }
    }
  });
}
