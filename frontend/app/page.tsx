import EntryForm from "@/components/base/entryForm";

export default function Home() {
    return (
        <main>
            <div className="flex justify-center items-center w-full h-screen md:pb-20">
                <div className="w-2/5">
                    <div className="mb-3">
                        <h1 className="text-xl font-bold">
                            Hey! 👋
                        </h1>
                    </div>
                    <EntryForm />
                </div>
            </div>
        </main>
    )
}
