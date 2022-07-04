import express from 'express';
import routes from './api_routes/routes.js';

const app = express();
routes(app);

const port = process.env.PORT || 3000;

app.listen(port, () => {
  console.log(`Listening to port http://localhost:${port}`)
})
