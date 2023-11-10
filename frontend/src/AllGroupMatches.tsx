import { Component, For, createMemo } from "solid-js";
import SingleGroupMatches from "./SingleGroupMatches";

const AllGroupMatches: Component<{ groups: any[] }> = (
  props
) => {
  return (
    <section>
      <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25">Group phase</h1>
      <div class="flex flex-row flex-wrap">
        <For each={props.groups}>
          {(group, i) => (
            <div class="p-4 w-1/2 xl:w-1/3">
              <SingleGroupMatches matches={group} />
            </div>
          )}
        </For>
      </div>
    </section>
  );
};

export default AllGroupMatches;
