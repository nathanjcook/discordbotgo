const model = require('../models/models.js');
const help = require('./help.js')

// add media
const getHelp =  async (req, res) => {
    try {
        res.status(200).json(help);
    } catch (err) {
        res.status(500).json([{"Error": err.message},]);
    }
};


module.exports = { getHelp}