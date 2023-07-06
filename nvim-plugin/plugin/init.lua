local api = vim.api
local uv = vim.loop

local function syncer()

    local child_cmd = "syncing"
    local child_input = uv.new_pipe()
    local child_output = uv.new_pipe()
    local child_err = uv.new_pipe()
    local options = {args = {"client","1"}, stdio = {child_input, child_output, child_err}}
    local handle;

    local on_exit = function (status)
        print(status)
        uv.close(handle)
    end

    handle = uv.spawn(child_cmd, options, on_exit)

    uv.read_start(child_output, function (status, data)
        if data then
            print(data)
        end
    end)

    uv.read_start(child_err, function (status, data)
        if data then
            print(data)
        end
    end)

    uv.write(child_input, "hello ")



end

syncer();
