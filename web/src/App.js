import { useState } from 'react'

function App() {
  const [rData, setRData] = useState("")
  const random = () => {
    fetch("http://127.0.0.1:8080/api/seed", {
      method: 'GET',
      dataType: 'json'
    }).then(res => res.json())
      .then(data => {
        let status = data.status
        if (status === 1) {
          setRData(data.data)
        } else {
          setRData("")
        }
      })
      .catch(
        () => {
          setRData("")
        }
      )
  }
  return (
    <div className="App">
      <div><button onClick={random}>Random</button></div>
      <div><input readOnly value={rData} /></div>
    </div>
  );
}

export default App;
