import { A } from "@solidjs/router";
import { Component, For, JSX, Show, createSignal } from "solid-js";
import {
  GroupInStore,
  Strength,
  createGroupsEmptyMatches,
  groups,
  groupIndex,
  matches,
  setGroups,
  setGroupIndex,
  setMatches,
} from "./groupStore";
import CreateGroup from "./CreateGroup";
import CreateGroupMatches from "./CreateGroupMatches";

const TournamentCreationStages = {
  Groups: 1,
  Matches: 2,
} as const;

type StageValues =
  (typeof TournamentCreationStages)[keyof typeof TournamentCreationStages];

const TournamentCustomGroups: Component = () => {
  const [groupName, setGroupName] = createSignal("");

  const [TournamentStage, setTournamentStage] = createSignal<StageValues>(
    TournamentCreationStages.Groups
  );

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
    const newSymbol = Symbol();
    const group: GroupInStore = {
      groupName: groupName(),
      countries: [
        { name: "", strength: Strength.Weak },
        { name: "", strength: Strength.Weak },
        { name: "", strength: Strength.Weak },
        { name: "", strength: Strength.Weak },
      ],
      index: newSymbol,
    };
    setGroups([...groups, group]);

    setGroupIndex([...groupIndex, newSymbol]);

    const thisGroupsMatches = createGroupsEmptyMatches(newSymbol);
    setMatches([...matches, ...thisGroupsMatches]);

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
          <p class="py-4">Create custom Tournament</p>
        </div>
      </header>
      <div class="p-8 flex">
        <div class="relative">
          <Show when={groupIndex.length < 6}>
            <label class="p-1 absolute bg-white left-2 -top-3 text-sm">
              Create new Group
            </label>
          </Show>
          <Show when={groupIndex.length >= 6}>
            <label class="p-1 absolute bg-white left-2 -top-3 text-sm">
              All Groups created
            </label>
          </Show>
          <input
            name="group-name"
            value={groupName()}
            class="p-4 mr-2 outline-1 border-2 rounded-lg outline-slate-200 focus:outline-slate-400 shadow"
            onInput={onGroupNameInput}
            onKeyDown={createNewGroupOnEnter}
            disabled={groupIndex.length >= 6}
          ></input>
        </div>
        <Show when={groupIndex.length < 6}>
          <button
            class="shadow p-4 rounded-md bg-slate-200 hover:bg-slate-300"
            onClick={createNewGroup}
            disabled={groupIndex.length >= 6}
          >
            Add
          </button>
        </Show>
        <Show when={groupIndex.length < 6}>
          <p class="mx-4 text-sm text-slate-400">
            {" "}
            {6 - groupIndex.length} Groups more required{" "}
          </p>
        </Show>

        <Show when={TournamentStage() == TournamentCreationStages.Groups}>
          <button
            class="shadow p-4 rounded-md bg-slate-200 hover:bg-slate-300"
            onClick={() => setTournamentStage(TournamentCreationStages.Matches)}
            disabled={groupIndex.length < 0}
            data-cy="manage-matches-button"
          >
            Manage Matches
          </button>
        </Show>

        <Show when={TournamentStage() == TournamentCreationStages.Matches}>
          <button
            class="shadow p-4 rounded-md bg-slate-200 hover:bg-slate-300"
            onClick={() => setTournamentStage(TournamentCreationStages.Groups)}
          >
            Back to Groups
          </button>
        </Show>
      </div>

      <Show when={TournamentStage() == TournamentCreationStages.Groups}>
        <div class="flex flex-wrap">
          <For each={groupIndex}>
            {(groupIndex) => <CreateGroup groupIndex={groupIndex} />}
          </For>
        </div>
      </Show>

      <Show when={TournamentStage() == TournamentCreationStages.Matches}>
        <div class="flex flex-wrap">
          <For each={groupIndex}>
            {(groupIndex) => <CreateGroupMatches groupIndex={groupIndex} />}
          </For>
        </div>
      </Show>

      <div>{JSON.stringify(groups)}</div>
      <hr />
      <div>{JSON.stringify(matches)}</div>
    </div>
  );
};

export default TournamentCustomGroups;
