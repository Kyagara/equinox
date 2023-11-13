const fs = require('fs')
const path = require('path')
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

const apiTemplateFiles = fs
  .readdirSync('./templates/api')
  .filter((file) => file.endsWith('.go.dt'))
  .map((file) => path.join('./templates/api', file))

const clients = ['lol', 'tft', 'lor', 'val', 'riot']

apiTemplateFiles.forEach((file) => {
  const fileName = path.basename(file, '.dt')
  compile(file, '../api', fileName)
})

clients.forEach((clientName) => {
  defs.clientName = clientName
  const name = clientName === 'riot' ? 'Riot' : clientName.toUpperCase()
  defs.clientNormalizedName = name

  const clientTemplateFiles = fs
    .readdirSync('./templates/clients')
    .filter((file) => file.endsWith('.go.dt'))
    .filter(
      (file) => !(clientName === 'riot' || clientName === 'lor') || !file.includes('constants.go'),
    )
    .map((file) => path.join('./templates/clients', file))

  clientTemplateFiles.forEach((file) => {
    const fileName = path.basename(file, '.dt')
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
  const pathName = path.dirname(outputPath)
  fs.mkdirSync(pathName, { recursive: true })
  fs.writeFileSync(outputPath, buffer, 'utf8')
}
