import { Component, For } from "solid-js";

const GroupMatches: Component = (props: {matches: Array<any>}) => {

  return (
    <section>
      <h1>Group phase</h1>
      <For each={props.matches}>{(match, i) =>
        <div class="flex">
          <div>{match.team1.name} - {match.team2.name} : {match.goalsTeam1} : {match.goalsTeam2}</div>
        </div>
        }
      </For>

    </section>
  )
}

export default GroupMatches
