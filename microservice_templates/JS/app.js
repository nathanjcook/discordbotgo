const express = require("express");
const dotenv = require('dotenv');

const app = express();
// Get config variables
dotenv.config();

const routes = require('./routes/routes');
const port = process.env.PORT | 3002;

app.use(express.json());
app.use(express.urlencoded({
    extended: true,
    })
);

// API Routes
app.use('/api', routes);

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
