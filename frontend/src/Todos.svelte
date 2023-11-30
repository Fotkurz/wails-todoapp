<script lang="ts">
    import Todo from "./Todo.svelte";
    import {RemoveTodo} from '../wailsjs/go/main/App.js'
    export let todos: string[] = []

    function removeTodo(i: number): void {
        console.log(`removing todo: ${i}`)
        RemoveTodo(i).then(result => todos = result)
    }
</script>
<!-- TODO: should find a way to invoke removeMe or get the index after button click -->
<div class="todo-list">
    <ul class="todos">
        {#each todos as todo, i}
            <Todo text={todo} index={i} />
        {/each}
    </ul>
</div>

<style>
    .todo-list {
        display: flex;
        flex-direction: column;
        align-items: center;
        border: 1px solid green;
    }

    .todos {
        list-style: none;
        border: 1px solid red;
        width: 400px;
        display: flex;
        flex-direction: column;
    }
</style>