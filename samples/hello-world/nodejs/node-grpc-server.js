express = require('express'),
app = express()

var port = process.env.PORT

if( !port ) {
  console.log( 'No port specified!' )
} else {
  app.listen(port, () => {
    console.log( 'Listening on', port )
  })
}
