//import * as https from "https";
const https = require("https");

const handler = async () => {
  const url = "";

  https
    .get(url, (res) => {
      let data = "";

      // データを受け取るたびにコールバックが呼ばれます
      res.on("data", (chunk) => {
        data += chunk;
      });

      // データの受信が完了したらコールバックが呼ばれます
      res.on("end", () => {
        console.log(data);
      });
    })
    .on("error", (err) => {
      console.error("Error:", err.message);
    });
};
handler();
