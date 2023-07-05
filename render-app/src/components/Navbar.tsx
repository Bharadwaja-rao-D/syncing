

const Navbar = () => {

    const start_collaboration = () => {
        console.debug("start collaboration")
    }

    return (
            <div className="navbar">
                <button onClick={start_collaboration}> colloborate </button>
            </div>
    );
}

export default Navbar;
