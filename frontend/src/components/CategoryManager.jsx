import React, { useState } from 'react';
import { Button, Drawer, List, Typography, Popconfirm, Form, Input, Space } from 'antd';
import { DeleteOutlined, PlusOutlined } from '@ant-design/icons';
import { useTodo } from '../context/TodoContext';

const { Text } = Typography;

const CategoryManager = () => {
  const [visible, setVisible] = useState(false);
  const { categories, createCategory, deleteCategory } = useTodo();
  const [form] = Form.useForm();

  const handleAdd = async () => {
    try {
      const values = await form.validateFields();
      await createCategory(values);
      form.resetFields();
    } catch (error) {
      console.error('Validation failed:', error);
    }
  };

  return (
    <>
      <Button onClick={() => setVisible(true)}>Manage Categories</Button>
      <Drawer
        title="Manage Categories"
        placement="right"
        onClose={() => setVisible(false)}
        open={visible}
      >
        <Form form={form} layout="inline" onFinish={handleAdd} style={{ marginBottom: 20 }}>
          <Form.Item
            name="name"
            rules={[{ required: true, message: 'Name is required' }]}
          >
            <Input placeholder="Category Name" />
          </Form.Item>
          <Form.Item name="color" initialValue="#3B82F6">
            <Input type="color" style={{ width: 50, padding: 0 }} />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" icon={<PlusOutlined />}>
              Add
            </Button>
          </Form.Item>
        </Form>

        <List
          dataSource={categories}
          renderItem={(item) => (
            <List.Item
              actions={[
                <Popconfirm
                  title="Delete category?"
                  description="Todos with this category will lose it."
                  onConfirm={() => deleteCategory(item.id)}
                  okText="Yes"
                  cancelText="No"
                >
                  <Button type="text" danger icon={<DeleteOutlined />} />
                </Popconfirm>,
              ]}
            >
              <Space>
                <div style={{ width: 16, height: 16, borderRadius: '50%', backgroundColor: item.color }} />
                <Text>{item.name}</Text>
              </Space>
            </List.Item>
          )}
        />
      </Drawer>
    </>
  );
};

export default CategoryManager;
