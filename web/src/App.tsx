import React, {useState} from 'react'

function App() {
    const [count, setCount] = useState(0)

    return (
        <div>
            <header>
                <p>Hello Audacia!</p>
                <p>
                    <button type="button" onClick={() => setCount((count) => count + 1)}>
                        count is: {count}
                    </button>
                </p>
            </header>
        </div>
    )
}

export default App
