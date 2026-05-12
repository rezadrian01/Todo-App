import React from 'react';
import { Input, Select, Space, Button } from 'antd';
import { SearchOutlined, ClearOutlined } from '@ant-design/icons';
import { useTodo } from '../context/TodoContext';

const { Option } = Select;

const TodoFilters = () => {
  const { filters, setFilters, categories, setPagination } = useTodo();

  const handleFilterChange = (key, value) => {
    setFilters(prev => ({ ...prev, [key]: value }));
    setPagination(prev => ({ ...prev, current: 1 })); // Reset page on filter change
  };

  const handleClearFilters = () => {
    setFilters({ search: '', status: '', category_id: '', priority: '' });
    setPagination(prev => ({ ...prev, current: 1 }));
  };

  return (
    <Space wrap style={{ marginBottom: 16 }}>
      <Input
        placeholder="Search tasks..."
        value={filters.search}
        onChange={(e) => handleFilterChange('search', e.target.value)}
        prefix={<SearchOutlined />}
        style={{ width: 200 }}
        allowClear
      />

      <Select
        placeholder="Status"
        value={filters.status || undefined}
        onChange={(val) => handleFilterChange('status', val)}
        style={{ width: 120 }}
        allowClear
      >
        <Option value="completed">Completed</Option>
        <Option value="incomplete">Incomplete</Option>
      </Select>

      <Select
        placeholder="Category"
        value={filters.category_id || undefined}
        onChange={(val) => handleFilterChange('category_id', val)}
        style={{ width: 150 }}
        allowClear
      >
        {categories.map(c => (
          <Option key={c.id} value={c.id}>{c.name}</Option>
        ))}
      </Select>

      <Select
        placeholder="Priority"
        value={filters.priority || undefined}
        onChange={(val) => handleFilterChange('priority', val)}
        style={{ width: 120 }}
        allowClear
      >
        <Option value="high">High</Option>
        <Option value="medium">Medium</Option>
        <Option value="low">Low</Option>
      </Select>

      <Button icon={<ClearOutlined />} onClick={handleClearFilters}>
        Clear
      </Button>
    </Space>
  );
};

export default TodoFilters;
