import React, { useState } from 'react';
import { Table, Tag, Popconfirm, Button, Space, Typography } from 'antd';
import { EditOutlined, DeleteOutlined, CheckCircleOutlined, SyncOutlined } from '@ant-design/icons';
import dayjs from 'dayjs';
import { useTodo } from '../context/TodoContext';
import TodoForm from './TodoForm';

const { Text } = Typography;

const priorityColors = {
  high: 'red',
  medium: 'orange',
  low: 'green',
};

const TodoList = () => {
  const { todos, categories, pagination, setPagination, loading, deleteTodo, toggleComplete } = useTodo();
  const [editingTodo, setEditingTodo] = useState(null);
  const [isModalVisible, setIsModalVisible] = useState(false);

  const getCategoryColor = (categoryId) => {
    const cat = categories.find(c => c.id === categoryId);
    return cat ? cat.color : 'default';
  };

  const getCategoryName = (categoryId) => {
    const cat = categories.find(c => c.id === categoryId);
    return cat ? cat.name : 'No Category';
  };

  const columns = [
    {
      title: 'Title',
      dataIndex: 'title',
      key: 'title',
      render: (text, record) => (
        <Text delete={record.completed}>{text}</Text>
      ),
    },
    {
      title: 'Category',
      dataIndex: 'category_id',
      key: 'category_id',
      render: (categoryId) => (
        categoryId ? (
          <Tag color={getCategoryColor(categoryId)}>{getCategoryName(categoryId)}</Tag>
        ) : '-'
      ),
    },
    {
      title: 'Priority',
      dataIndex: 'priority',
      key: 'priority',
      render: (priority) => (
        <Tag color={priorityColors[priority] || 'default'}>{priority.toUpperCase()}</Tag>
      ),
    },
    {
      title: 'Due Date',
      dataIndex: 'due_date',
      key: 'due_date',
      render: (date) => (date ? dayjs(date).format('MMM D, YYYY') : '-'),
    },
    {
      title: 'Status',
      key: 'status',
      dataIndex: 'completed',
      render: (completed) => (
        <Tag icon={completed ? <CheckCircleOutlined /> : <SyncOutlined spin />} color={completed ? 'success' : 'processing'}>
          {completed ? 'DONE' : 'IN PROGRESS'}
        </Tag>
      ),
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button 
            type="text" 
            icon={record.completed ? <SyncOutlined /> : <CheckCircleOutlined />} 
            onClick={() => toggleComplete(record.id)}
          />
          <Button 
            type="text" 
            icon={<EditOutlined />} 
            onClick={() => {
              setEditingTodo(record);
              setIsModalVisible(true);
            }} 
          />
          <Popconfirm
            title="Delete the task"
            description="Are you sure to delete this task?"
            onConfirm={() => deleteTodo(record.id)}
            okText="Yes"
            cancelText="No"
          >
            <Button type="text" danger icon={<DeleteOutlined />} />
          </Popconfirm>
        </Space>
      ),
    },
  ];

  const handleTableChange = (newPagination) => {
    setPagination({
      ...pagination,
      current: newPagination.current,
      pageSize: newPagination.pageSize,
    });
  };

  return (
    <>
      <Table 
        columns={columns} 
        dataSource={todos} 
        rowKey="id" 
        pagination={{
          current: pagination.current,
          pageSize: pagination.pageSize,
          total: pagination.total,
          showSizeChanger: true,
        }}
        loading={loading}
        onChange={handleTableChange}
        scroll={{ x: 800 }}
      />
      <TodoForm 
        visible={isModalVisible} 
        onClose={() => {
          setIsModalVisible(false);
          setEditingTodo(null);
        }} 
        initialValues={editingTodo} 
      />
    </>
  );
};

export default TodoList;
