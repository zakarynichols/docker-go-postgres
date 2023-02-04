import { dirname } from "path";
import { fileURLToPath } from "url";
import express from "express";
import { join } from "path";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const app = express();
const port = 7777;

app.use(express.static(join(__dirname, "../public")));

app.get("/", (_req, res) => {
  res.sendFile(join(__dirname, "public", "index.html"));
});

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});