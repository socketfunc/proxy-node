const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const glob = require('glob')

const PROTO_PATHS = glob.sync(`${__dirname}/../proto/**/*.proto`)
const protoDef = protoLoader.loadSync(PROTO_PATHS, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
})
const proto = grpc.loadPackageDefinition(protoDef)

const host = 'localhost:5000'
const client = new proto.gateway.GatewayService(host, grpc.credentials.createInsecure())

const conn = client.subscribe({
  id: 'req-id',
  topic: 'req-topic'
})

conn.on('data', (data) => {
  console.log(data)
})
