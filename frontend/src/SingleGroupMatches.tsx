import { Component, For } from "solid-js";
import TeamMatchGroup from "./TeamMatchGroup";

const SingleGroupMatches: Component<{ matches: any[] }> = (props) => {
  return (
    <div>
      <h2 class="pb-2 text-lg"><u>{props.matches[0].groupName}</u></h2>
      <div class="grid grid-cols-[auto_auto_auto_auto] max-w-fit">
      <For each={props.matches}>
        {(match, i) => (
          <TeamMatchGroup match={match} />
        )}
      </For>
      </div>
    </div>
  );
};

export default SingleGroupMatches;
