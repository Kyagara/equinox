const fs = require('fs')
const doT = require('dot')
const glob = require('glob')
process.chdir(__dirname)
global.require = require
const defs = {}

doT.templateSettings = {
  evaluate: /\r?\n?\{\{([\s\S]+?)\}\}/g,
  interpolate: /\r?\n?\{\{=([\s\S]+?)\}\}/g,
  encode: /\r?\n?\{\{!([\s\S]+?)\}\}/g,
  use: /\r?\n?\{\{#([\s\S]+?)\}\}/g,
  define: /\r?\n?\{\{##\s*([\w.$]+)\s*(:|=)([\s\S]+?)#\}\}/g,
  conditional: /\r?\n?\{\{\?(\?)?\s*([\s\S]*?)\s*\}\}/g,
  iterate: /\r?\n?\{\{~\s*(?:\}\}|([\s\S]+?)\s*:\s*([\w$]+)\s*(?::\s*([\w$]+))?\s*\}\})/g,
  varname: 'it',
  strip: false,
  append: false,
  selfcontained: false,
}

glob.sync('./templates/api/*.go.dt').forEach((file) => {
  const fileName = file.split('/')[3].replace('.dt', '')
  compile(file, '../api', fileName)
})

const clients = ['lol', 'tft', 'lor', 'val', 'riot']

clients.forEach((clientName) => {
  defs.clientName = clientName
  let name = clientName.toUpperCase()
  name = name === 'RIOT' ? 'Riot' : name
  defs.clientNormalizedName = name

  glob.sync('./templates/clients/*.go.dt').forEach((file) => {
    if ((clientName === 'riot' || clientName === 'lor') && file.includes('constants.go')) return
    const fileName = file.split('/')[3].replace('.dt', '')
    compile(file, `../clients/${clientName}`, fileName)
  })
})

function compile(inputPath, outputPath, fileName) {
  const input = readFile(inputPath)
  const output = compileTemplate(input, inputPath)
  const path = `${outputPath}/${fileName}`

  saveTemplate(output, path)
}

function readFile(path) {
  return fs.readFileSync(path, 'utf8')
}

function compileTemplate(input, fileName) {
  console.log(`Compiling '${fileName}'.`)
  try {
    return doT.template(input, undefined, defs)({})
  } catch (err) {
    throw new Error({ message: `Error compiling '${fileName}'`, error: err })
  }
}

function saveTemplate(buffer, outputPath) {
  console.log(`Writing '${outputPath}'.`)
  const pathName = `../${outputPath.replace(/^\.*\/|\/?[^/]+\.[a-z]+|\/$/g, '')}`
  if (!fs.existsSync(pathName)) {
    fs.mkdirSync(pathName, { recursive: true })
  }
  fs.writeFileSync(outputPath, buffer, 'utf8')
}
