/* @refresh reload */
import { render } from "solid-js/web";
import "solid-devtools";

import "./index.css";
import Tournament from "./pages/Tournament";
import { Router, Routes, Route } from "@solidjs/router";
import Home from "./pages/Home";
import TournamentCustomGroups from "./pages/TournamentCustomGroups";

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
        <Route path="/" component={Home}/>
        {/* <Route path="/new" component={TournamentCustomGroups}></Route> */}
        <Route path="/2021" component={Tournament("http://localhost:3000/api/2021", "EM soccer tournament simulator 2021")}/>
        <Route path="/2024/" component={Tournament("http://localhost:3000/api/2024", "EM soccer tournament simulator 2024")}/>
      </Routes>
    </Router>
  ),
  root!
);
