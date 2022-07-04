'use strict';
import segments from '../service/strava.js'

let controllers = {
  segments: (req, res) => {
    segments.find(req, res, (err, dist) => {
      if (err)
        res.send(err);
      res.json(dist);
    })
  }
};

export default controllers
