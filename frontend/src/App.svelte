<script lang="ts">
  import {AddTodo} from '../wailsjs/go/main/App.js'
  import Todos from './Todos.svelte';

  let allTodos: string[] = ["Hello", "World", "!"]
  let newTodo: string = ''
  
  function addTodo(): void {
    
    AddTodo(newTodo).then(result => allTodos = result)
    newTodo = ''
  }
</script>

<main>
  <h1>Wails Todo APP</h1>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={newTodo} class="input" id="newTodo" type="text"/>
    <button class="btn" on:click={addTodo}>+</button>
  </div>
  
  <Todos todos={allTodos} />
</main>

<style>

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
