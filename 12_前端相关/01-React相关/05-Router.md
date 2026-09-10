### 安装
  - `npm i react-router-dom`

### 代码案例-创建路由配置
```jsx
// src/router/index.js
import { createBrowserRouter } from 'react-router-dom';
import Layout from '../components/Layout';
import Home from '../pages/Home';
import About from '../pages/About';
import User from '../pages/User';
import NotFound from '../pages/NotFound'; // 导入 404 组件

// 方式一：直接配置路由数组
export const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        index: true, // 默认子路由，访问 / 时显示
        element: <Home />,
      },
      {
        path: 'about',
        element: <About />,
      },
      {
        path: 'user/:id', // 动态路由参数
        element: <User />,
      },
      // ⚠️ 重要：放在 children 最后，匹配该层级下所有未定义的路径
      { 
        path: '*', 
        element: <NotFound />, 
      },
    ],
  },
]);
```

### 代码案例-在入口文件使用 RouterProvider
```jsx
// src/main.jsx 或 index.js
import React from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';
import { router } from './router';

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
```

### 代码案例-申明式导航
```jsx
// src/components/Layout.jsx
import { Outlet, Link, NavLink } from 'react-router-dom';

function Layout() {
  return (
    <div>
      <header style={{ padding: '20px', background: '#f0f0f0' }}>
        <h1>我的网站</h1>
        <nav>
          {/* Link：跳转链接 申明式导航 */}
          <Link to="/" style={{ marginRight: '10px' }}>首页</Link>
          <Link to="/about" style={{ marginRight: '10px' }}>关于</Link>
          <Link to="/user/123" style={{ marginRight: '10px' }}>用户123</Link>
          
          {/* NavLink：带高亮效果的链接（activeClassName 或 className 函数） */}
          <NavLink 
            to="/" 
            style={({ isActive }) => ({
              fontWeight: isActive ? 'bold' : 'normal',
              color: isActive ? 'red' : 'blue',
              marginRight: '10px'
            })}
          >
            首页(NavLink)
          </NavLink>
        </nav>
      </header>
      
      {/* Outlet：子路由组件渲染的位置 */}
      <main style={{ padding: '20px' }}>
        <Outlet />
      </main>
      
      <footer style={{ padding: '10px', background: '#f0f0f0' }}>
        © 2026 我的网站
      </footer>
    </div>
  );
}

export default Layout;
```

### 代码案例-二级路由页面组件
```jsx
// src/pages/Home.jsx
function Home() {
  return (
    <div>
      <h2>🏠 首页</h2>
      <p>欢迎来到我的网站！</p>
    </div>
  );
}

export default Home;
```

### 代码案例-二级路由页面组件
```jsx
// src/pages/About.jsx
function About() {
  return (
    <div>
      <h2>📖 关于我们</h2>
      <p>这是一个使用 React Router v6 的示例项目。</p>
    </div>
  );
}

export default About;
```

### 代码案例-编程式导航、动态路由参数
```jsx
// src/pages/User.jsx
import { useLocation, useParams, useSearchParams, useNavigate } from 'react-router-dom';

function User() {
  // 1. 获取当前路径信息
  const location = useLocation();
  console.log('当前路径：', location.pathname);
  console.log('传递的 state：', location.state);
  
  // 2.useParams：获取动态路由参数（如 /user/:id）
  const { id } = useParams();

  // 3. 获取查询参数（如 ?name=张三&age=18）
  const [searchParams, setSearchParams] = useSearchParams();
  const name = searchParams.get('name');
  const age = searchParams.get('age');
  
  // useNavigate：编程式导航
  const navigate = useNavigate();

  const goBack = () => {
    navigate(-1); // 返回上一页
  };

  const goHome = () => {
    navigate('/'); // 跳转到首页
  };
    
  const goAbout = () => {
      navigate('/about', { state: { from: 'user' } })
  }
  return (
    <div>
      <h2>👤 用户详情</h2>
      <p>当前用户 ID：<strong>{id}</strong></p>
      <button onClick={goBack}>返回</button>
      <button onClick={goHome} style={{ marginLeft: '10px' }}>首页</button>
      <button onClick={goAbout} style={{ marginLeft: '10px' }}>首页</button>
    </div>
  );
}

export default User;
```

### 代码案例-404页面组件
```jsx
// src/pages/NotFound.jsx
import { Link } from 'react-router-dom';

function NotFound() {
  return (
    <div style={{ 
      textAlign: 'center', 
      padding: '60px 20px',
      fontFamily: 'Arial'
    }}>
      <h1 style={{ fontSize: '100px', margin: '0', color: '#e74c3c' }}>404</h1>
      <h2 style={{ color: '#333' }}>页面未找到</h2>
      <p style={{ color: '#666', marginBottom: '20px' }}>
        抱歉，您访问的页面不存在或已被移除。
      </p>
      <Link 
        to="/" 
        style={{
          display: 'inline-block',
          padding: '10px 24px',
          background: '#3498db',
          color: '#fff',
          textDecoration: 'none',
          borderRadius: '4px'
        }}
      >
        返回首页
      </Link>
    </div>
  );
}

export default NotFound;
```

### 申明式导航
  - `<Link to="/"></Link>`标签

### 编程式导航
  - `useNavigate`的`Hook`函数

### 传参
