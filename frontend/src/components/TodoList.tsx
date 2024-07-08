import React from 'react';
import { Todo } from '../types';

type TodoListProps = {
  todos: Todo[];
  updateTodo: (ID: number, completed: boolean) => void;
  deleteTodo: (ID: number) => void;
};

const TodoList: React.FC<TodoListProps> = ({ todos, updateTodo, deleteTodo }) => {
  return (
    <ul>
      {todos.map((todo) => (
        <li key={todo.ID}>
          <span
            style={{ textDecoration: todo.completed ? 'line-through' : 'none' }}
            onClick={() => updateTodo(todo.ID, !todo.completed)}
          >
            {todo.title}
          </span>
          <button onClick={() => updateTodo(todo.ID, todo.completed)}>Update</button>
          <button onClick={() => deleteTodo(todo.ID)}>Delete</button>
        </li>
      ))}
    </ul>
  );
};

export default TodoList;
