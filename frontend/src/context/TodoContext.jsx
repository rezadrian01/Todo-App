import React, { createContext, useContext, useState, useEffect, useCallback } from 'react';
import { message } from 'antd';
import * as api from '../services/api';

const TodoContext = createContext();

export function TodoProvider({ children }) {
  const [todos, setTodos] = useState([]);
  const [categories, setCategories] = useState([]);
  const [pagination, setPagination] = useState({ current: 1, pageSize: 10, total: 0 });
  const [filters, setFilters] = useState({
    search: '', status: '', category_id: '', priority: ''
  });
  const [loading, setLoading] = useState(false);

  const fetchCategories = useCallback(async () => {
    try {
      const res = await api.getCategories();
      setCategories(res.data.data);
    } catch (error) {
      message.error('Failed to load categories');
    }
  }, []);

  const fetchTodos = useCallback(async () => {
    setLoading(true);
    try {
      const params = { page: pagination.current, limit: pagination.pageSize, ...filters };
      // Remove empty filters
      Object.keys(params).forEach(key => {
        if (params[key] === '' || params[key] === null || params[key] === undefined) {
          delete params[key];
        }
      });

      const res = await api.getTodos(params);
      setTodos(res.data.data);
      if (res.data.pagination) {
        setPagination(p => ({ ...p, total: res.data.pagination.total }));
      }
    } catch (error) {
      message.error('Failed to load todos');
    } finally {
      setLoading(false);
    }
  }, [pagination.current, pagination.pageSize, filters]);

  useEffect(() => {
    fetchCategories();
  }, [fetchCategories]);

  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  // Actions
  const handleCreateTodo = async (values) => {
    await api.createTodo(values);
    message.success('Todo created');
    fetchTodos();
  };

  const handleUpdateTodo = async (id, values) => {
    await api.updateTodo(id, values);
    message.success('Todo updated');
    fetchTodos();
  };

  const handleDeleteTodo = async (id) => {
    await api.deleteTodo(id);
    message.success('Todo deleted');
    // If it's the last item on the page, go to previous page (unless it's page 1)
    if (todos.length === 1 && pagination.current > 1) {
      setPagination(p => ({ ...p, current: p.current - 1 }));
    } else {
      fetchTodos();
    }
  };

  const handleToggleComplete = async (id) => {
    await api.toggleTodoComplete(id);
    message.success('Status updated');
    fetchTodos();
  };

  const handleCreateCategory = async (values) => {
    await api.createCategory(values);
    message.success('Category created');
    fetchCategories();
  };

  const handleDeleteCategory = async (id) => {
    await api.deleteCategory(id);
    message.success('Category deleted');
    fetchCategories();
    fetchTodos(); // Refresh todos as some might have lost their category
  };

  const contextValue = {
    todos,
    categories,
    pagination,
    setPagination,
    filters,
    setFilters,
    loading,
    createTodo: handleCreateTodo,
    updateTodo: handleUpdateTodo,
    deleteTodo: handleDeleteTodo,
    toggleComplete: handleToggleComplete,
    createCategory: handleCreateCategory,
    deleteCategory: handleDeleteCategory,
    fetchTodos
  };

  return (
    <TodoContext.Provider value={contextValue}>
      {children}
    </TodoContext.Provider>
  );
}

export const useTodo = () => useContext(TodoContext);
