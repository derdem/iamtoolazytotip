import { A } from "@solidjs/router";
import { Component, For, JSX, createSignal } from "solid-js";
import { groups, setGroups } from "./groupStore";

const TournamentCustom: Component = () => {

  const [groupName, setGroupName] = createSignal("");
  const onGroupNameInput: JSX.EventHandler<HTMLInputElement, InputEvent> = (event) => {
    const groupName = event.currentTarget.value;
    setGroupName(groupName);
  }
  const createNewGroup = () => {
    setGroups([...groups, groupName()]);
  }
  const createNewGroupOnEnter: JSX.EventHandler<HTMLInputElement, KeyboardEvent> = (event) => {
    if (event.key === "Enter") {
      createNewGroup();
    }
  }

  return (
    <div>
      <header class="bg-sky-800 text-center text-white flex justify-between items-center">
        <div class="flex text-2xl ml-4">
          <A href="/" class="no-underline">
            <i class="py-4 mr-4 fa-solid fa-house"></i>
          </A>
          <p class="py-4">EM soccer tournament simulator 2024</p>
        </div>
      </header>
      <div class="p-8">
        <input class="p-4 mr-2 outline-1 border-2 rounded-lg outline-slate-200 focus:outline-slate-400" onInput={onGroupNameInput} onKeyDown={createNewGroupOnEnter}></input>
        <button class="p-4 rounded-md bg-slate-200 hover:bg-slate-300" onClick={createNewGroup}>Create new Group</button>
      </div>
      <div>
        <For each={groups}>
          {(group) => <div class="p-4">{group}</div>}
        </For>
      </div>
    </div>
  );
};

export default TournamentCustom;
