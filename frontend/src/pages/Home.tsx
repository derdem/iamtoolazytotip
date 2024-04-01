import { Component } from "solid-js";
import EnterTournamentCard from "../EnterTournamentCard";

const Home: Component = () => {
  return (
    <div>
      <header class="bg-sky-800 text-center text-white flex justify-between items-center">
        <div class="text-2xl ml-4">
          <p class="py-4">Tournament Overview</p>
        </div>
      </header>
      <div class="container mx-auto">
        <div class="flex justify-between p-8">
          <EnterTournamentCard name="New Tournament" url="/new" />
          <EnterTournamentCard name="Tournament 2021" url="/2021" />
          <EnterTournamentCard name="Tournament 2024" url="/2024" />
        </div>
      </div>
    </div>
  );
};

export default Home;
