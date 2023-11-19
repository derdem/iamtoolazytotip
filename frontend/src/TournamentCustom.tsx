import { A } from "@solidjs/router";
import { Component, For, JSX, createSignal } from "solid-js";
import { GroupInStore, Strength, groups, setGroups } from "./groupStore";
import CreateGroup from "./CreateGroup";

const TournamentCustom: Component = () => {
  const [groupName, setGroupName] = createSignal("");
  const [groupIndex, setGroupIndex] = createSignal<number[]>([]);
  const onGroupNameInput: JSX.EventHandler<HTMLInputElement, InputEvent> = (
    event
  ) => {
    const groupName = event.currentTarget.value;
    setGroupName(groupName);
  };
  const createNewGroup = () => {
    if (groupName() === "") {
      return;
    }
    const group: GroupInStore = {
      groupName: groupName(),
      countries: [
        { name: "", strength: Strength.Weak },
        { name: "", strength: Strength.Weak },
        { name: "", strength: Strength.Weak },
        { name: "", strength: Strength.Weak },
      ],
      index: groups.length,
    };
    setGroups([...groups, group]);

    const newGroupIndex = groupIndex().length;
    setGroupIndex([...groupIndex(), newGroupIndex])
    setGroupName("");
  };
  const createNewGroupOnEnter: JSX.EventHandler<
    HTMLInputElement,
    KeyboardEvent
  > = (event) => {
    if (event.key === "Enter") {
      createNewGroup();
    }
  };

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
      <div class="p-8 flex">
        <div class="relative">
          <label class="p-1 absolute bg-white left-2 -top-3 text-sm">
            Create new Group
          </label>
          <input
            name="group-name"
            value={groupName()}
            class="p-4 mr-2 outline-1 border-2 rounded-lg outline-slate-200 focus:outline-slate-400 shadow"
            onInput={onGroupNameInput}
            onKeyDown={createNewGroupOnEnter}
          ></input>
        </div>
        <button
          class="shadow p-4 rounded-md bg-slate-200 hover:bg-slate-300"
          onClick={createNewGroup}
        >
          Add
        </button>
      </div>
      <div class="flex">
        <For each={groupIndex()}>{(groupIndex) => <CreateGroup groupIndex={groupIndex} />}</For>
      </div>
      <div>{JSON.stringify(groups)}</div>
    </div>
  );
};

export default TournamentCustom;
