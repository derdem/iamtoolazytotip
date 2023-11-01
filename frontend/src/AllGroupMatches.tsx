import { Component, For } from "solid-js";
import SingleGroupMatches from "./SingleGroupMatches";

const AllGroupMatches: Component<{ groups: {[key: string]: any[]} }> = (props) => {
  const groups = Object.values(props.groups)
  console.log(groups)

  return (
    <section>
      <h1>Group phase</h1>
      <For each={groups}>
        {(group, i) => (
          <SingleGroupMatches matches={group} />
        )}
      </For>
    </section>
  );
};

export default AllGroupMatches;
