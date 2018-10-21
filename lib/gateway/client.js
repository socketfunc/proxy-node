const events = require('events')

const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const glob = require('glob')

const PROTO_PATHS = glob.sync(`${__dirname}/../proto/**/*.proto`)
const protoDef = protoLoader.loadSync(PROTO_PATHS, {})
const proto = grpc.loadPackageDefinition(protoDef)

const host = 'k8s-master:30107'
const client = new proto.gateway.GatewayService(host, grpc.credentials.createInsecure())

const conn = client.subscribe({
  id: 'req-id',
  topic: 'req-topic'
})

conn.on('data', (data) => {
  console.log(data)
})

class GatewayClient extends events.EventEmitter {
  constructor (host) {
    super()
    this.client = new proto.gateway.GatewayService(host, grpc.credentials.createInsecure())
  }

  publish (topic, payload) {

  }

  subscribe (topic, id) {

  }

  unsubscribe (topic, id) {

  }
}
