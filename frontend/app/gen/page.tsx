import ChatResponse from "@/components/base/chatResponse";

export default function Gen(){
    return (
        <main>
            <div className="flex justify-center items-center w-full h-screen md:pb-20">
                <div className="w-2/5">
                    <div className="mb-3">
                        <h1 className="text-xl font-bold">
                            Talk
                        </h1>
                        <ChatResponse prompt="Hello! Who are you?" />
                    </div>
                </div>
            </div>
        </main>
    )
}
