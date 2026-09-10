### Redux管理数据流
 - `Action` -> `Reducer` -> `Store(很多个state)`
 -  `Action` <- `View(视图)` <- `Store(很多个state)`

#### 三个核心概念
  - `state`: 一个对象，存放着我们管理的数据状态
  - `action`: 一个对象，用来描述你想怎么改数据
  - `reducer`: 一个函数，根据`action`的描述生成一个新的`state`

### 插件
  - `Redux Toolkit`: 官方推荐编写`Redux`逻辑的方式，是一套工具的集合集，简化书写方式
    - 简化`store`的配置方式
    - 内置`immer`支持可变式状态修改
    - 内置`thunk`更好的异步创建
  - `react-redux`: 用来链接`Redux`和`React组件`的中间件
      -           --获取状态-->
      - `redux`  `react-redux` `React组件`
      -           <--更新状态--
  - `npx create-react-app react-redux`
  - `npm i @reduxjs/toolkit react-redux`
  - `npm install @reduxjs/toolkit react-redux`

### 目录
  - 通常集中状态管理的部分都会单独创建一个单独的`store`目录
  - 应用通常会有很多个子`store`模块，所以会创建一个`modules`目录，在内部编写业务分类的子`store`
  - `store`中的入口文件`index.js`的作用是组合`modules`中的所有子模块，并导出`store`

### 工程使用
  - 组件中使用`useSelector`函数获取`store`中的数据
  - 组件中使用`useDispatch`函数获取`dispatch`方法
  - 通过执行`store`模块中导出的`actionCreater`方法提交`action`对象

### 代码案例-创建Redux State Slice
```js
// src/store/modules/counterSlice.js
import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  value: 0,
};

export const counterSlice = createSlice({
  name: 'counter', // 这个 slice 的名字，会作为 action type 的前缀
  initialState,
  reducers: {
    // 这里定义的函数就是 "reducer"
    // 注意：可以直接写 "修改" 逻辑，createSlice 内部用 Immer 库保证了数据的不可变性
    increment: (state) => {
      state.value += 1;
    },
    decrement: (state) => {
      state.value -= 1;
    },
    incrementByAmount: (state, action) => {
      state.value += action.payload;
    },
  },
});

// 自动生成的 action creators，导出供组件使用
export const { increment, decrement, incrementByAmount } = counterSlice.actions;

// 导出的 reducer，用来更新 store 中的对应状态
export default counterSlice.reducer;
```

### 代码案例-将Slice挂载到Store
```js
// src/store/index.js
import { configureStore } from '@reduxjs/toolkit';
import counterReducer from './modules/counterSlice';

export const store = configureStore({
  reducer: {
    counter: counterReducer, // 这样 state 中就会有一个 counter 字段
  },
});
```

### 代码案例-将Store提供给React
```js
// src/index.js (或 main.jsx)
import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'react-redux';
import { store } from './store';
import App from './App';

ReactDOM.createRoot(document.getElementById('root')).render(
  <Provider store={store}>
    <App />
  </Provider>
);
```

### 代码案例-在组件中使用Redux状态和操作
```jsx
// src/App.jsx
import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { increment, decrement } from './store/modules/counterSlice';

export function App() {
  // 1. 使用 useSelector 读取 store 中的数据
  const count = useSelector((state) => state.counter.value);
  // 2. 使用 useDispatch 获取 dispatch 函数
  const dispatch = useDispatch();

  return (
    <div>
      <button onClick={() => dispatch(increment())}>+</button>
      <span>{count}</span>
      <button onClick={() => dispatch(decrement())}>-</button>
    </div>
  );
}
```

### 全局状态-Context或Redux、Zustand
- 定义一个`reducer`函数(根据当前想要做的修改返回一个新的状态): 数据不可变，基于原始状态生成一个新的状态
- 使用`createStore`方法传入`reducer`函数生成一个`store`实例对象
- 使用`store`实例的`subscribe`方法订阅数据的变化(数据一旦变化，可以得到通知): 回调函数可以在每次state发生变化时自动执行
- 使用`store`实例的`dispatch`方法提交`action`对象触发数据变化(告诉`reducer`你想怎么改数据)
- 使用`store`实例的`getState`方法获取最新的状态数据更新到视图中