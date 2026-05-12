import React, { useState } from 'react';
import { Layout, Typography, Card, Space, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import TodoList from './components/TodoList';
import TodoFilters from './components/TodoFilters';
import CategoryManager from './components/CategoryManager';
import TodoForm from './components/TodoForm';

const { Header, Content } = Layout;
const { Title } = Typography;

const AppContent = () => {
  const [isNewModalVisible, setIsNewModalVisible] = useState(false);

  return (
    <Layout style={{ minHeight: '100vh', backgroundColor: '#f0f2f5' }}>
      <Header style={{ background: '#fff', padding: '0 24px', display: 'flex', alignItems: 'center' }}>
        <Title level={3} style={{ margin: 0 }}>Industrix Todo</Title>
      </Header>
      <Content style={{ padding: '24px', maxWidth: 1200, margin: '0 auto', width: '100%' }}>
        <Card>
          <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 24, flexWrap: 'wrap', gap: 16 }}>
            <TodoFilters />
            <Space>
              <CategoryManager />
              <Button type="primary" icon={<PlusOutlined />} onClick={() => setIsNewModalVisible(true)}>
                New Task
              </Button>
            </Space>
          </div>
          <TodoList />
        </Card>
      </Content>
      <TodoForm 
        visible={isNewModalVisible} 
        onClose={() => setIsNewModalVisible(false)} 
      />
    </Layout>
  );
};

export default AppContent;
