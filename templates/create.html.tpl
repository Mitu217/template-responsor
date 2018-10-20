<!DOCTYPE html>
<html>
<body>
    <div>
        <form action="/admin/create" method="post">
            <p>Name：
                <input type="text" name="name" size="40">
            </p>
            <p>AccessType：
                <input type="radio" name="request-method" value="0">GET
                <input type="radio" name="request-method" value="1">POST
            </p>
            <p>EncodingType：
                <select name="content-type">
                    <option value="0">application/json</option>
                    <option value="1">application/msgpack</option>
                </select>
            </p>
            <p>Route：
                <input type="text" name="route" size="40">
            </p>
            <p>Data：<br>
                <textarea name="data" rows="4" cols="40"></textarea>
            </p>
            <p>
                <input type="submit" value="Create"><input type="reset" value="Reset">
            </p>
        </form>
    </div>
</body>
</html>