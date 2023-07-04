import "./globals.css";
import classnames from "classnames";
import { Inter } from "next/font/google";

export const metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={classnames("bg-slate-950 font-inter text-slate-400")}>
        {children}
      </body>
    </html>
  );
}
