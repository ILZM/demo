import { useNavigate } from "react-router-dom";
import { signOut } from "supertokens-auth-react/recipe/session";
import { recipeDetails } from "../config";
import CallAPIView from "./CallAPIView";
import { FaBlog } from "react-icons/fa";

interface ILink {
    name: string;
    onClick: () => void;
}

export default function SuccessView(props: { userId: string }) {
    let userId = props.userId;

    const navigate = useNavigate();

    async function logoutClicked() {
        await signOut();
        navigate("/auth");
    }

    function openLink(url: string) {
        window.open(url, "_blank");
    }

    const links: ILink[] = [
        {
            name: "Blogs",
            onClick: () => openLink("https://supertokens.com/blog")
        },
        {
            name: "Documentation",
            onClick: () => openLink(recipeDetails.docsLink)
        },
        {
            name: "Sign Out",
            onClick: logoutClicked
        },
    ];

    return (
        <>
            <div className="main-container">
                <div className="top-band success-title bold-500">
                    <FaBlog></FaBlog>
                </div>
                <div className="inner-content">
                    <div>Your userID is:</div>
                    <div className="truncate" id="user-id">
                        {userId}
                    </div>
                    <CallAPIView />
                </div>
            </div>
            <div className="bottom-links-container">
                {links.map((link) => (
                    <div className="link" key={link.name}>
                        <FaBlog></FaBlog>
                        <div role={"button"} onClick={link.onClick}>
                            {link.name}
                        </div>
                    </div>
                ))}
            </div>
            <FaBlog></FaBlog>
        </>
    );
}
