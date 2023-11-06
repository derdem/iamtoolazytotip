/* @refresh reload */
import { render } from "solid-js/web";
import "solid-devtools";

import "./index.css";
import Tournament2021 from "./Tournament2021";
import Tournament2024 from "./Tournament2024";
import { Router, Routes, Route } from "@solidjs/router";

const root = document.getElementById("root");

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    "Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got mispelled?"
  );
}

render(
  () => (
    <Router>
      <Routes>
        <Route path="/" component={Tournament2021}/>
        <Route path="/2024/" component={Tournament2024}/>
      </Routes>
    </Router>
  ),
  root!
);
