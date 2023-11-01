import { Component, createSignal, onMount } from "solid-js";

import styles from "./App.module.css";
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
  console.log(sortedMatches)
  return sortedMatches;
}

const App: Component = () => {
  const [tournamentOutcome, setTournamentOutcome] = createSignal<{
    group: any[];
    sixteen: any[];
    eight: any[];
    four: any[];
    final: any[];
  }>({ group: [], sixteen: [], eight: [], four: [], final: [] });
const [groupOutcomes, setGroupCoutcomes] = createSignal<{[key: string]: any[];}>({});


  onMount(async () => {
    const response = await fetch("http://localhost:8080/api/");
    const data = await response.json();
    data.final = [data.final];
    console.log(data);
    console.log(data.final[0].team1.name);
    setTournamentOutcome(data);
    setGroupCoutcomes(sortGroupMatches(data.group));
  });

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <p class="pt-8">EM soccer tournament simulator</p>
      </header>
      <AllGroupMatches groups={groupOutcomes()} />
      <KoRound matches={tournamentOutcome().sixteen} name="Round of 16" />
      <KoRound matches={tournamentOutcome().eight} name="Round of 8" />
      <KoRound matches={tournamentOutcome().four} name="Round of 4" />
      <KoRound matches={tournamentOutcome().final} name="Final" />
      {/* tournamentOutcome().final */}
    </div>
  );
};

export default App;
