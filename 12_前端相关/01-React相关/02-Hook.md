### 逻辑运算符
  - `isLogin && <span>已经登录</span>`
  - `isLogin ? <span>已经登录</span> : <span>请登录</span>`


## 几种类型的Hook
  - 状态管理
    - `useState`:
    - `useReducer`:
  - 副作用处理
    - `useEffect`:
    - `useLayoutEffect`:
  - 性能优化
    - `useMemo`:
    - `useCallback`:
  - 上下文
    - `useContext`
  - DOM/Ref
    - `useRef`:
    - `useImperativeHandle`:

### useEffect
  - 函数组件中执行"副作用"
  - 副作用: 不属于组件渲染流程的操作
    - 数据请求
    - DOM操作
    - 订阅、监听
    - 定时器
    - 手动修改DOM
  - 使用方式:
    ```tsx
        useEffect(() => {
            // 副作用代码
            console.log('组件渲染了');
            
            // 清理函数
            return () => {
                console.log('组件要卸载了，清理资源');
            }   
        }, [])
    ```



### useMemo
  - 缓存函数执行后的结果，节省 CPU
    - 使用场景:
      ```tsx
        function ServerList({ servers, keyword }) {
            // ❌ 每次渲染都重新过滤
            // 如果 servers 有 10000 条，每次渲染都要遍历一遍，很浪费
            const filtered = servers.filter(s => s.ip.includes(keyword));
    
            // ✅ 只有 servers 或 keyword 变化时才重新过滤
            const filtered2 = useMemo(() => {
                console.log('重新过滤了！'); // 依赖不变时，这行不会打印
                return servers.filter(s => s.ip.includes(keyword));
            }, [servers, keyword]);
            
            return (
                <Table dataSource={filtered2} />
            );
        }
      ```

### useCallback
  - 缓存函数，避免组件渲染时重复创建函数

  - 使用场景:
    - 传给子组件
       ```tsx
        const Child = React.memo(({ onClick }) => {
            // 只有 props 变化时才重新渲染
        });
        
        function Parent() {
            const handleClick = useCallback(() => {
                console.log('点击');
            }, []);
            
            return <Child onClick={handleClick} />;
        }
        ```
    - 作为useEffect的依赖
      ```tsx
        function Search({ keyword }) {
            const doSearch = useCallback(() => {
                fetch(`/api/search?q=${keyword}`);
            }, [keyword]);
        
            useEffect(() => {
                doSearch(); // 用了 doSearch
            }, [doSearch]); // ← 如果不用 useCallback，doSearch 每次都变，useEffect 会反复执行
        }
        ```
      
## Hook的参数和外部变量
  - `参数`: 来自组件作用域，须写在依赖数组里。
  - `外部变量`: 在函数调用时才传入的，不依赖组件的状态变化。

## 依赖数组
  - 不看顺序