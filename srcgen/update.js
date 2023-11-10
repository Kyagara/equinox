const fs = require('fs')
const fetch = require('node-fetch-commonjs')
process.chdir(__dirname)

const files = [
  ['http://www.mingweisamuel.com/riotapi-schema/openapi-3.0.0.json', './specs/.spec.json'],
  ['http://www.mingweisamuel.com/riotapi-schema/enums/queues.json', './specs/.queues.json'],
  ['http://www.mingweisamuel.com/riotapi-schema/enums/queueTypes.json', './specs/.queueTypes.json'],
  ['http://www.mingweisamuel.com/riotapi-schema/enums/gameTypes.json', './specs/.gameTypes.json'],
  ['http://www.mingweisamuel.com/riotapi-schema/enums/gameModes.json', './specs/.gameModes.json'],
  ['http://www.mingweisamuel.com/riotapi-schema/enums/maps.json', './specs/.maps.json'],
  ['http://www.mingweisamuel.com/riotapi-schema/routesTable.json', './specs/.routesTable.json'],
]

if (!fs.existsSync('./specs')) {
  fs.mkdirSync('./specs')
}

files.forEach(async (file) => {
  console.log(`Downloading '${file[1]}'.`)
  let res = await fetch(file[0])
  let data = await res.json()
  console.log(`Writing '${file[1]}'.`)
  fs.writeFileSync(file[1], JSON.stringify(data))
})
