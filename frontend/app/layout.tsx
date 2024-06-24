import type { Metadata } from "next";
import { Libre_Baskerville } from "next/font/google";
import "./globals.css";

const libre_baskerville = Libre_Baskerville({ weight: ["400", "700"], style: "normal", subsets: ["latin"] });

export const metadata: Metadata = {
    title: "MediKiosk",
    description: "Medical Advice to everyone",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body className={libre_baskerville.className}>{children}</body>
        </html>
    );
}
