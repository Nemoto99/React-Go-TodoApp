import React, { useEffect, useState } from 'react';
import { Todo } from './types';
import axios from 'axios';
import TodoList from './components/TodoList';

const App: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState<string>('');

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const response = await axios.get('http://localhost:8080/todos');
    setTodos(response.data);
  };

  const addTodo = async () => {
    await axios.post('http://localhost:8080/todos', { title: newTodo });
    fetchTodos();
    setNewTodo('');
  };

  const updateTodo = async (ID: number, completed: boolean) => {
    try {
      await axios.put<Todo>(`http://localhost:8080/todos/${ID}`, {
        completed: !completed,
      });
      fetchTodos();
    } catch (error) {
      console.error('Error updating todo:', error);
    }
  };

  const deleteTodo = async (ID: number) => {
    try {
      await axios.delete(`http://localhost:8080/todos/${ID}`);
      fetchTodos();
    } catch (error) {
      console.error('Error deleting todo:', error);
    }
  };

  return (
    <div>
      <h1>Todo App</h1>
      <input
        type='text'
        value={newTodo}
        onChange={(e) => setNewTodo(e.target.value)}
      />
      <button onClick={addTodo}>Add</button>
      <TodoList
        todos={todos}
        updateTodo={updateTodo}
        deleteTodo={deleteTodo}
      />
    </div>
  );
};

export default App;
