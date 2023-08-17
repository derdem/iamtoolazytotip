import { Component, createSignal, onMount } from 'solid-js';

import styles from './App.module.css';
import GroupMatches from './GroupMatches';
import KoRound from './KoRound';

const App: Component = () => {

  const [tournamentOutcome, setTournamentOutcome] = createSignal({})
  onMount(async () => {
    const response = await fetch("http://localhost:8080/api/")
    const data = await response.json()
    data.final = [data.final]
    console.log(data)
    console.log(data.final[0].team1.name)
    setTournamentOutcome(data)
  })

  return (
    <div class={styles.App}>
      <header class={styles.header}>

        <p>
          EM soccer tournament simulator
        </p>
      </header>
      <GroupMatches matches={tournamentOutcome().group}/>
      <KoRound matches={tournamentOutcome().sixteen} name="Round of 16" />
      <KoRound matches={tournamentOutcome().eight} name="Round of 8" />
      <KoRound matches={tournamentOutcome().four} name="Round of 4" />
      <KoRound matches={tournamentOutcome().final} name="Final" />
      {/* tournamentOutcome().final */}
    </div>
  );
};

export default App;
