import { Component, For, createMemo } from "solid-js";
import TeamMatchGroup from "./TeamMatchGroup";

const SingleGroupMatches: Component<{ matchResults: any[], groupName: string }> = (props) => {
  // const comparePlaytime = (a: any, b: any) => {
  //   if (a.playtime < b.playtime) {
  //     return -1;
  //   }
  //   if (a.playtime > b.playtime) {
  //     return 1;
  //   }
  //   return 0;
  // }
  //const matches = createMemo(() => props.matches.sort(comparePlaytime));

  return (
    <div>
      <h2 class="pb-2 text-lg"><u>{props.groupName}</u></h2>
      <div class="grid grid-cols-[auto_auto_auto_auto] max-w-fit">
      <For each={props.matchResults}>
        {(matchResult, i) => (
          <TeamMatchGroup matchResult={matchResult} />
        )}
      </For>
      </div>
    </div>
  );
};

export default SingleGroupMatches;
