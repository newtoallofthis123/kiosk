"use client"

import ChatResponse from "@/components/base/chatResponse";
import { useSearchParams } from "next/navigation"

export default function Probs() {
    const searchParams = useSearchParams()
    const problems = searchParams.get('problems') || '';

    const problemList = problems.split(";")
    const disease = searchParams.get('disease') || 'Common Cold';

    const prompt = `Imagine you are a doctor and you are treating a 
patient with the following problems: ${problemList.join(", ")}.
Name is ${searchParams.get('name') || 'Anonymous'}.
It has been determined that the patient has been diagnosed with the following: ${disease}.
Now, explain it to the patient in a way that they can understand.
`;

    return (
        <main>
            <div className="flex justify-center items-center w-full md:pt-20">
                <div className="w-2/5">
                    <div className="mb-3">
                        <h1 className="text-2xl font-bold">
                            Problems
                        </h1>
                        <p className="text-sm">
                            The following problems have been identified in your text:
                        </p>
                        <ul className="border-2 border-black rounded-md p-2 mt-3">
                            {problemList.map((problem, index) => (
                                <li key={index}>{problem}</li>
                            ))}
                        </ul>
                        <div className="mt-2">
                            <ChatResponse prompt={prompt} />
                        </div>
                    </div>
                </div>
            </div>
        </main>
    )
}
