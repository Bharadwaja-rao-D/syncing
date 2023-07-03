import Pallet from "./components/Pallet"

function App() {

    const pitems = [{icon: "rect", name: "rectangle"},{icon: "ci", name: "circle"}]

  return (
    <>
     <Pallet pitems={pitems} />
    </>
  )
}

export default App
