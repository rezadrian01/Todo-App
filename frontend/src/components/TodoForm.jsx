import React, { useEffect } from 'react';
import { Modal, Form, Input, Select, DatePicker } from 'antd';
import dayjs from 'dayjs';
import { useTodo } from '../context/TodoContext';

const { Option } = Select;

const TodoForm = ({ visible, onClose, initialValues }) => {
  const [form] = Form.useForm();
  const { createTodo, updateTodo, categories } = useTodo();

  useEffect(() => {
    if (visible) {
      if (initialValues) {
        form.setFieldsValue({
          ...initialValues,
          due_date: initialValues.due_date ? dayjs(initialValues.due_date) : null,
          category_id: initialValues.category_id || undefined,
        });
      } else {
        form.resetFields();
      }
    }
  }, [visible, initialValues, form]);

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      const payload = {
        ...values,
        due_date: values.due_date ? values.due_date.toISOString() : null,
      };

      if (initialValues) {
        await updateTodo(initialValues.id, payload);
      } else {
        await createTodo(payload);
      }
      onClose();
    } catch (error) {
      console.error('Validation failed:', error);
    }
  };

  return (
    <Modal
      title={initialValues ? 'Edit Task' : 'New Task'}
      open={visible}
      onOk={handleSubmit}
      onCancel={onClose}
      afterClose={() => form.resetFields()}
      destroyOnClose
    >
      <Form form={form} layout="vertical" initialValues={{ priority: 'medium' }}>
        <Form.Item
          name="title"
          label="Title"
          rules={[{ required: true, message: 'Please input the title!' }]}
        >
          <Input placeholder="What needs to be done?" />
        </Form.Item>

        <Form.Item name="description" label="Description">
          <Input.TextArea rows={3} placeholder="Add more details..." />
        </Form.Item>

        <Form.Item name="category_id" label="Category">
          <Select placeholder="Select a category" allowClear>
            {categories.map((cat) => (
              <Option key={cat.id} value={cat.id}>
                {cat.name}
              </Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item name="priority" label="Priority">
          <Select>
            <Option value="high">High</Option>
            <Option value="medium">Medium</Option>
            <Option value="low">Low</Option>
          </Select>
        </Form.Item>

        <Form.Item name="due_date" label="Due Date">
          <DatePicker style={{ width: '100%' }} />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default TodoForm;
