import React from "react";
import {Email, GitHub} from "@mui/icons-material";

function Home() {
    return (
        <div className="home">
            <div className="about">
                <h2>Hi, My Name is Alex</h2>
                <div className="prompt">
                    <p>A software developer with a passion for learning and creating.</p>
                    <Email />
                    <GitHub />
                </div>
            </div>
            <div className="skills">
                <h1>Skills: </h1>
                <ol></ol>
            </div>
        </div>
    )
}

export default Home;