import React from "react";
import {GitHub, Instagram, Twitter} from "@mui/icons-material";
import '../styles/Footer.css';

function Footer() {
    return (
        <div className="footer">
            <div className="socialMedia">
                <Instagram />
                <Twitter />
                <GitHub />
            </div>
            <p> &copy; 2023 sw1pr0g.com</p>
        </div>
    )
}

export default Footer;