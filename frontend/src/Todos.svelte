<script lang="ts">
    import Todo from "./Todo.svelte";
    import {RemoveTodo} from '../wailsjs/go/main/App.js'
    import RemoveButton from "./RemoveButton.svelte";
    
    export let todos: string[] = []

    function removeTodo(i: number): void {
        console.log(`removing todo: ${i}`)
        RemoveTodo(i).then(result => todos = result)
    }
</script>
<!-- TODO: should find a way to invoke removeMe or get the index after button click -->
<div class="todo-container">
    <ul class="todos">
        {#each todos as todo, i}
            <div class="todo-item">
                <Todo text={todo} />
                <RemoveButton on:click={() => removeTodo(i)} />
            </div>
        {/each}
    </ul>
</div>

<style>
    .todo-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        /* border: 1px solid green; */
    }

    .todos {
        list-style: none;
        /* border: 1px solid red; */
        width: 400px;
        padding-left: 0px;
        flex-direction: column;
        align-items: center;
    }

    .todo-item {
        background: lightcyan;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        min-height: 25px;
        width: 100%;
        border-radius: 10px;
        /* border: 1px solid lightcoral; */
        margin-top: 10px;
    }

</style>