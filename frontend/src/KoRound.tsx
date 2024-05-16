import { Component, For, Show } from "solid-js";
import TeamMatchKo from "./TeamMatchKo";

const KoRound: Component<{ matchResults: any[]; name: string }> = (props) => {
  console.log(props.matchResults)
  return (
    <section class="p-4 xl:w-1/3 w-1/2">
      <h1 class="pb-2 text-lg"><u>{props.name}</u></h1>
      <div class="grid grid-cols-[auto_auto_auto_auto] max-w-fit">
        <For each={props.matchResults}>
          {(matchResult, i) => (
            <TeamMatchKo matchResult={matchResult}/>
          )}
        </For>
      </div>
    </section>
  );
};

export default KoRound;
