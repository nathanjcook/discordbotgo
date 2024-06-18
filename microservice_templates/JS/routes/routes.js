const express = require('express');

const controller = require('../controllers/controllers');

const router = express.Router();

// Route paths for api calls
router.post('/help', controller.getHelp);

module.exports = router;