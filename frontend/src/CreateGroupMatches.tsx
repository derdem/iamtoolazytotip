import { Component, For } from "solid-js";
import { groups } from "./groupStore";

interface CreateGroupProps {
  groupIndex: number;
}

const CreateGroupMatches: Component<CreateGroupProps> = (props) => {
  const groupName = groups[props.groupIndex].groupName;
  const countries = groups[props.groupIndex].countries;

  return (
    <div class="m-4 p-2 border shadow">
      <h1 class="underline mb-2">{groupName}</h1>
      <For each={[0, 1, 2, 3, 4, 5]}>
        {(matchIndex) => (
          <div class="flex">
            <div class="m-4">
              <select
                name={groupName + "-" + (matchIndex + 1).toString() + "-1"}
                id={groupName + "-" + (matchIndex + 1).toString() + "-1"}
              >
                <For each={countries}>
                  {(country) => (
                    <option value={country.name}>{country.name}</option>
                  )}
                </For>
              </select>
            </div>
            <div class="pt-4"> - </div>
            <div class="m-4">
              <select
                name={groupName + "-" + (matchIndex + 1).toString() + "-2"}
                id={groupName + "-" + (matchIndex + 1).toString() + "-2"}
              >
                <For each={countries}>
                  {(country) => (
                    <option value={country.name}>{country.name}</option>
                  )}
                </For>
              </select>
            </div>
          </div>
        )}
      </For>
    </div>
  );
}

export default CreateGroupMatches;
