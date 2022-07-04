'use strict';
import controller from '../controllers/controller.js';

export default (app) => {
  app.route('/routes').get(controller.segments)
}
