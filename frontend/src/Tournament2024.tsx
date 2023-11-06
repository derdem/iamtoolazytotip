import { Component, Show, createEffect, createSignal, onMount } from "solid-js";

import AllGroupMatches from "./AllGroupMatches";
import KoRound from "./KoRound";

const sortGroupMatches = (groupMatches: any[]) => {
  const sortedMatches: any = {};
  groupMatches.forEach((match) => {
    if (!sortedMatches[match.groupNumber]) {
      sortedMatches[match.groupNumber] = [];
    }
    sortedMatches[match.groupNumber].push(match);
  });
  console.log(sortedMatches);
  return sortedMatches;
};

const getNewSimulation = async () => {
  const response = await fetch("http://localhost:8080/api/2024");
  const data = await response.json();
  data.final = [data.final];
  return data;
}

const App: Component = () => {
  const [tournamentOutcome, setTournamentOutcome] = createSignal<{
    group: any[];
    sixteen: any[];
    eight: any[];
    four: any[];
    final: any[];
  }>({ group: [], sixteen: [], eight: [], four: [], final: [] });
  const [groupOutcomes, setGroupOutcomes] = createSignal<{
    [key: string]: any[];
  }>({});
  createEffect(() =>
    console.log("The latest groupOutcomes are", groupOutcomes())
  );

  const newSimulation = async () => {
    const data = await getNewSimulation();
    setTournamentOutcome(data);
    setGroupOutcomes(sortGroupMatches(data.group));
  }

  onMount(async () => {
    newSimulation()
  });

  return (
    <div>
      <header class="bg-sky-800 text-center text-white flex justify-between items-center">
        <div class="text-2xl ml-4">
          <p class="py-4">EM soccer tournament simulator 2024</p>
        </div>

        <button onClick={newSimulation} class="bg-white text-black rounded-md px-4 py-2 mr-4">Run simulation</button>
      </header>
      <AllGroupMatches groups={groupOutcomes()} />
      <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25">KO phase</h1>
      <div class="flex flex-row flex-wrap">
        <KoRound matches={tournamentOutcome().sixteen} name="Round of 16" />
        <KoRound matches={tournamentOutcome().eight} name="Round of 8" />
        <KoRound matches={tournamentOutcome().four} name="Round of 4" />
      </div>
      <div class="flex justify-around">
        <KoRound matches={tournamentOutcome().final} name="Final" />
      </div>
      <Show when={tournamentOutcome().final.length > 0} fallback={<h1 class="p-4 text-xl bg-sky-800 bg-opacity-25">No Winner yet</h1>}>
        <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25 text-center">Winner: {tournamentOutcome().final[0].winner.name}</h1>
      </Show>

      {/* tournamentOutcome().final */}
    </div>
  );
};

export default App;
