const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const glob = require('glob')

const PROTO_PATHS = glob.sync(`${__dirname}/../proto/**/*.proto`)
const protoDef = protoLoader.loadSync(PROTO_PATHS, {
  keepCase: true,
  longs: String,
  defaults: true,
  oneofs: true
})
const proto = grpc.loadPackageDefinition(protoDef)

const server = new grpc.Server()

function run () {
  server.addService(proto.gateway.GatewayService.service, {
    publish: (call, callback) => {
      console.log(call.request)
      callback(null, {})
    },
    subscribe: async (conn) => {
      console.log(conn.request)
      conn.write({
        packet: {
          id: 'req-id',
          topic: 'req-topic',
          event: 'req-event',
          payload: Buffer.from('req-payload')
        }
      })
    },
    unsubscribe: (call, callback) => {
      callback(null, {})
    },
    ping: (conn) => {
      conn.write({})
    }
  })
  server.bind('localhost:5000', grpc.ServerCredentials.createInsecure())
  server.start()
}

run()
